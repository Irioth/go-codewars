package runner

import (
	"codewars"
	"errors"
	"time"
)

type MessageType byte

const (
	_ MessageType = iota
	GameOver
	AuthenticationToken
	TeamSize
	ProtocolVersion
	GameContext
	PlayerContext
	Move
)

const Version int = 1

/**
 * Стратегия --- интерфейс, содержащий описание методов искусственного интеллекта армии.
 * Каждая пользовательская стратегия должна реализовывать этот интерфейс.
 * Может отсутствовать в некоторых языковых пакетах, если язык не поддерживает интерфейсы.
 */
type Strategy interface {

	/**
		 * Основной метод стратегии, осуществляющий управление армией. Вызывается каждый тик.
		 *
	     * me    Информация о вашем игроке.
	     * world Текущее состояние мира.
	     * game  Различные игровые константы.
	     * move  Результатом работы метода является изменение полей данного объекта.
	*/
	Move(*codewars.Player, *codewars.World, *codewars.Game, *codewars.Move)
}

const DefaultHost = "127.0.0.1"
const DefaultPort = "31001"
const DefaultToken = "0000000000000000"

type CodeWars struct {
	*AiCup

	players    map[int64]*codewars.Player
	facilities map[int64]*codewars.Facility

	terrainByCellXY [][]codewars.TerrainType
	weatherByCellXY [][]codewars.WeatherType
}

func Start(s Strategy, args ...string) {
	var host, port, token string

	if len(args) == 3 {
		host, port, token = args[0], args[1], args[2]
	} else {
		host, port, token = DefaultHost, DefaultPort, DefaultToken
	}

	cli := &CodeWars{
		AiCup:      new(AiCup),
		players:    make(map[int64]*codewars.Player),
		facilities: make(map[int64]*codewars.Facility),
	}

	for {
		if err := cli.Dial(host, port); err == nil {
			cli.writeToken(token)
			cli.writeProtoVersion(Version)
			cli.ReadTeamSize()

			g := cli.readGame()

			for {
				p, w, err := cli.readContext()
				if err != ErrGameOver {
					m := new(codewars.Move)
					m.VehicleType = codewars.Vehicle_None
					m.Action = codewars.Action_None

					s.Move(p, w, g, m)
					cli.writeMove(m)
				}
			}

			cli.Close()
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func (c *CodeWars) readGame() *codewars.Game {
	c.ensureMessageType(GameContext)

	if c.readBool() {
		return &codewars.Game{
			RandomSeed:                             c.readInt64(),
			TickCount:                              c.readInt(),
			WorldWidth:                             c.readFloat64(),
			WorldHeight:                            c.readFloat64(),
			FogOfWarEnabled:                        c.readBool(),
			VictoryScore:                           c.readInt(),
			FacilityCaptureScore:                   c.readInt(),
			VehicleEliminationScore:                c.readInt(),
			ActionDetectionInterval:                c.readInt(),
			BaseActionCount:                        c.readInt(),
			AdditionalActionCountPerControlCenter:  c.readInt(),
			MaxUnitGroup:                           c.readInt(),
			TerrainWeatherMapColumnCount:           c.readInt(),
			TerrainWeatherMapRowCount:              c.readInt(),
			PlainTerrainVisionFactor:               c.readFloat64(),
			PlainTerrainStealthFactor:              c.readFloat64(),
			PlainTerrainSpeedFactor:                c.readFloat64(),
			SwampTerrainVisionFactor:               c.readFloat64(),
			SwampTerrainStealthFactor:              c.readFloat64(),
			SwampTerrainSpeedFactor:                c.readFloat64(),
			ForestTerrainVisionFactor:              c.readFloat64(),
			ForestTerrainStealthFactor:             c.readFloat64(),
			ForestTerrainSpeedFactor:               c.readFloat64(),
			ClearWeatherVisionFactor:               c.readFloat64(),
			ClearWeatherStealthFactor:              c.readFloat64(),
			ClearWeatherSpeedFactor:                c.readFloat64(),
			CloudWeatherVisionFactor:               c.readFloat64(),
			CloudWeatherStealthFactor:              c.readFloat64(),
			CloudWeatherSpeedFactor:                c.readFloat64(),
			RainWeatherVisionFactor:                c.readFloat64(),
			RainWeatherStealthFactor:               c.readFloat64(),
			RainWeatherSpeedFactor:                 c.readFloat64(),
			VehicleRadius:                          c.readFloat64(),
			TankDurability:                         c.readInt(),
			TankSpeed:                              c.readFloat64(),
			TankVisionRange:                        c.readFloat64(),
			TankGroundAttackRange:                  c.readFloat64(),
			TankAerialAttackRange:                  c.readFloat64(),
			TankGroundDamage:                       c.readInt(),
			TankAerialDamage:                       c.readInt(),
			TankGroundDefence:                      c.readInt(),
			TankAerialDefence:                      c.readInt(),
			TankAttackCooldownTicks:                c.readInt(),
			TankProductionCost:                     c.readInt(),
			IFVDurability:                          c.readInt(),
			IFVSpeed:                               c.readFloat64(),
			IFVVisionRange:                         c.readFloat64(),
			IFVGroundAttackRange:                   c.readFloat64(),
			IFVAerialAttackRange:                   c.readFloat64(),
			IFVGroundDamage:                        c.readInt(),
			IFVAerialDamage:                        c.readInt(),
			IFVGroundDefence:                       c.readInt(),
			IFVAerialDefence:                       c.readInt(),
			IFVAttackCooldownTicks:                 c.readInt(),
			IFVProductionCost:                      c.readInt(),
			ARRVDurability:                         c.readInt(),
			ARRVSpeed:                              c.readFloat64(),
			ARRVVisionRange:                        c.readFloat64(),
			ARRVGroundDefence:                      c.readInt(),
			ARRVAerialDefence:                      c.readInt(),
			ARRVProductionCost:                     c.readInt(),
			ARRVRepairRange:                        c.readFloat64(),
			ARRVRepairSpeed:                        c.readFloat64(),
			HelicopterDurability:                   c.readInt(),
			HelicopterSpeed:                        c.readFloat64(),
			HelicopterVisionRange:                  c.readFloat64(),
			HelicopterGroundAttackRange:            c.readFloat64(),
			HelicopterAerialAttackRange:            c.readFloat64(),
			HelicopterGroundDamage:                 c.readInt(),
			HelicopterAerialDamage:                 c.readInt(),
			HelicopterGroundDefence:                c.readInt(),
			HelicopterAerialDefence:                c.readInt(),
			HelicopterAttackCooldownTicks:          c.readInt(),
			HelicopterProductionCost:               c.readInt(),
			FighterDurability:                      c.readInt(),
			FighterSpeed:                           c.readFloat64(),
			FighterVisionRange:                     c.readFloat64(),
			FighterGroundAttackRange:               c.readFloat64(),
			FighterAerialAttackRange:               c.readFloat64(),
			FighterGroundDamage:                    c.readInt(),
			FighterAerialDamage:                    c.readInt(),
			FighterGroundDefence:                   c.readInt(),
			FighterAerialDefence:                   c.readInt(),
			FighterAttackCooldownTicks:             c.readInt(),
			FighterProductionCost:                  c.readInt(),
			MaxFacilityCapturePoints:               c.readFloat64(),
			FacilityCapturePointsPerVehiclePerTick: c.readFloat64(),
			FacilityWidth:                          c.readFloat64(),
			FacilityHeight:                         c.readFloat64(),
		}
	}

	return nil
}

var ErrGameOver = errors.New("game over")

func (c *CodeWars) readContext() (*codewars.Player, *codewars.World, error) {
	switch c.readOpcode() {
	case GameOver:
		return nil, nil, ErrGameOver
	case PlayerContext:
		if c.readBool() {
			return c.readPlayer(), c.readWorld(), nil
		} else {
			return nil, nil, nil
		}
	default:
		return nil, nil, ErrWrongType
	}
}

func (c *CodeWars) readPlayer() *codewars.Player {
	switch c.readByte() {
	case 0:
		return nil
	case 127:
		return c.players[c.readInt64()] // consume id, no changes
	default:
		p := &codewars.Player{
			Id:              c.readInt64(),
			Me:              c.readBool(),
			StrategyCrashed: c.readBool(),
			Score:           c.readInt(),
			RemainingActionCooldownTicks: c.readInt(),
		}

		c.players[p.Id] = p

		return p
	}
}

func (c *CodeWars) readWorld() *codewars.World {
	if c.readBool() {
		w := new(codewars.World)

		w.TickIndex = c.readInt()
		w.TickCount = c.readInt()
		w.Width = c.readFloat64()
		w.Height = c.readFloat64()

		c.readPlayers(w)
		c.readVehicles(w)       // New
		c.readVehiclesUpdate(w) // Updates

		if w.TickIndex == 0 {
			c.readTerrains(w)
			c.readWeather(w)
		}

		c.readFacilities(w)
	}
}

func (c *CodeWars) writeMove(m *runner.Move) {
	c.writeOpcode(Move)

	c.writeBool(true)

	c.writeByte(byte(m.Action))
	c.writeInt(m.Group)
	c.writeFloat64(m.Left)
	c.writeFloat64(m.Top)
	c.writeFloat64(m.Right)
	c.writeFloat64(m.Bottom)
	c.writeFloat64(m.X)
	c.writeFloat64(m.Y)
	c.writeFloat64(m.Angle)
	c.writeFloat64(m.MaxSpeed)
	c.writeFloat64(m.MaxAngularSpeed)
	c.writeByte(byte(m.VehicleType))
	c.writeInt64(m.FacilityId)

	c.flush()
}

func (c *CodeWars) readWeather() (weather [][]codewars.WeatherType) {
	for i := c.readInt(); i > 0; i-- {
		var slice []codewars.WeatherType
		for j := c.readInt(); j > 0; j-- {
			slice = append(slice, codewars.WeatherType(c.readByte()))
		}
		weather = append(weather, slice)
	}

	return
}

func (c *CodeWars) readTerrains() (terrains [][]codewars.TerrainType) {
	for i := c.readInt(); i > 0; i-- {
		var slice []codewars.TerrainType
		for j := c.readInt(); j > 0; j-- {
			slice = append(slice, codewars.TerrainType(c.readByte()))
		}
		terrains = append(terrains, slice)
	}
	return
}

func (c *CodeWars) readFacility(fn func(id int64) *runner.Facility) {
	switch c.readByte() {
	case 0, 127:
		break
	default:
		f := fn(c.readInt64())
		f.FacilityType = runner.FacilityType(c.readByte())
		f.OwnerPlayerId = c.readInt64()
		f.Left = c.readFloat64()
		f.Top = c.readFloat64()
		f.CapturePoints = c.readFloat64()
		f.VehicleType = runner.VehicleType(c.readByte())
		f.ProductionProgress = c.readInt()
	}
}

func (c *CodeWars) readVehicleUpdate(fn func(id int64) *runner.Vehicle) *runner.Vehicle {
	if c.readBool() {
		v := fn(c.readInt64())
		v.X = c.readFloat64()
		v.Y = c.readFloat64()
		v.Durability = c.readInt()
		v.RemainingAttackCooldownTicks = c.readInt()
		v.Selected = c.readBool()
		v.Groups = c.readIntArray()

		return v
	}

	return nil
}

func (c *CodeWars) readVehicle(fn func(id int64) *runner.Vehicle) {
	if c.readBool() {
		v := fn(c.readInt64())
		v.X = c.readFloat64()
		v.Y = c.readFloat64()
		v.Radius = c.readFloat64()
		v.PlayerId = c.readInt64()
		v.Durability = c.readInt()
		v.MaxDurability = c.readInt()
		v.MaxSpeed = c.readFloat64()
		v.VisionRange = c.readFloat64()
		v.SquaredVisionRange = c.readFloat64()
		v.GroundAttackRange = c.readFloat64()
		v.SquaredGroundAttackRange = c.readFloat64()
		v.AerialAttackRange = c.readFloat64()
		v.SquaredAerialAttackRange = c.readFloat64()
		v.GroundDamage = c.readInt()
		v.AerialDamage = c.readInt()
		v.GroundDefence = c.readInt()
		v.AerialDefence = c.readInt()
		v.AttackCooldownTicks = c.readInt()
		v.RemainingAttackCooldownTicks = c.readInt()
		v.Type = runner.VehicleType(c.readByte())
		v.Aerial = c.readBool()
		v.Selected = c.readBool()
		v.Groups = c.readIntArray()
	}
}

func (c *CodeWars) readVehiclesUpdate(w *runner.World) {
	w.Updates = w.Updates[:0]
	for l := c.readInt(); l > 0; l-- {
		if v := c.readVehicleUpdate(w.Vehicle); v != nil {
			w.Updates = append(w.Updates, v.Id)
		}
	}
}

func (c *CodeWars) readFacilities(w *runner.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readFacility(w.Facility)
	}
}

func (c *CodeWars) readVehicles(w *runner.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readVehicle(w.Vehicle)
	}
}

func (c *CodeWars) readPlayers(w *runner.World) {
	for l := c.readInt(); l > 0; l-- {
		c.readPlayer(w.Player)
	}
}