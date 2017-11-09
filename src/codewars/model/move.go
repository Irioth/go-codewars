package model

import "sync"

/**
 * Стратегия игрока может управлять юнитами посредством установки свойств объекта данного класса.
 */
type Move struct {
	/**
	 * Устанавливает действие игрока.
	 */
	Action ActionType
	/**
	 * Устанавливает группу юнитов для различных действий.
	 * <p>
	 * Является опциональным параметром для действий {@code ActionType.CLEAR_AND_SELECT},
	 * {@code ActionType.ADD_TO_SELECTION} и {@code ActionType.DESELECT}. Если для этих действий группа юнитов
	 * установлена, то параметр {@code VehicleType}, а также параметры прямоугольной рамки {@code Left}, {@code Top},
	 * {@code Right} и {@code Bottom} будут проигнорированы.
	 * <p>
	 * Является обязательным параметром для действий {@code ActionType.ASSIGN}, {@code ActionType.DISMISS} и
	 * {@code ActionType.DISBAND}. Для действия {@code ActionType.DISBAND} является единственным учитываемым параметром.
	 * <p>
	 * Корректными значениями являются целые числа от {@code 1} до {@code game.MaxUnitGroup} включительно.
	 */
	Group int
	/**
	 * @return Устанавливает левую границу прямоугольной рамки для выделения юнитов.
	 * <p>
	 * Является обязательным параметром для действий {@code ActionType.CLEAR_AND_SELECT},
	 * {@code ActionType.ADD_TO_SELECTION} и {@code ActionType.DESELECT}, если не установлена группа юнитов.
	 * В противном случае граница будет проигнорирована.
	 * <p>
	 * Корректными значениями являются вещественные числа от {@code 0.0} до {@code Right} включительно.
	 */
	Left float64
	/**
	 * @return Устанавливает верхнюю границу прямоугольной рамки для выделения юнитов.
	 * <p>
	 * Является обязательным параметром для действий {@code ActionType.CLEAR_AND_SELECT},
	 * {@code ActionType.ADD_TO_SELECTION} и {@code ActionType.DESELECT}, если не установлена группа юнитов.
	 * В противном случае граница будет проигнорирована.
	 * <p>
	 * Корректными значениями являются вещественные числа от {@code 0.0} до {@code Bottom} включительно.
	 */
	Top float64
	/**
	 * @return Устанавливает правую границу прямоугольной рамки для выделения юнитов.
	 * <p>
	 * Является обязательным параметром для действий {@code ActionType.CLEAR_AND_SELECT},
	 * {@code ActionType.ADD_TO_SELECTION} и {@code ActionType.DESELECT}, если не установлена группа юнитов.
	 * В противном случае граница будет проигнорирована.
	 * <p>
	 * Корректными значениями являются вещественные числа от {@code Left} до {@code game.WorldWidth} включительно.
	 */
	Right float64
	/**
	 * @return Устанавливает нижнюю границу прямоугольной рамки для выделения юнитов.
	 * <p>
	 * Является обязательным параметром для действий {@code ActionType.CLEAR_AND_SELECT},
	 * {@code ActionType.ADD_TO_SELECTION} и {@code ActionType.DESELECT}, если не установлена группа юнитов.
	 * В противном случае граница будет проигнорирована.
	 * <p>
	 * Корректными значениями являются вещественные числа от {@code Top} до {@code game.WorldHeight} включительно.
	 */
	Bottom float64
	/**
	 * Устанавливает абсциссу точки или вектора.
	 * <p>
	 * Является обязательным параметром для действия {@code ActionType.MOVE} и задаёт целевую величину смещения юнитов
	 * вдоль оси абсцисс.
	 * <p>
	 * Является обязательным параметром для действия {@code ActionType.ROTATE} и задаёт абсциссу точки, относительно
	 * которой необходимо совершить поворот.
	 * <p>
	 * Корректными значениями для действия {@code ActionType.MOVE} являются вещественные числа от
	 * {@code -game.WorldWidth} до {@code game.WorldWidth} включительно. Корректными значениями для действия
	 * {@code ActionType.ROTATE} являются вещественные числа от {@code -game.WorldWidth} до
	 * {@code 2.0 * game.WorldWidth} включительно.
	 */
	X float64
	/**
	 * Устанавливает ординату точки или вектора.
	 * <p>
	 * Является обязательным параметром для действия {@code ActionType.MOVE} и задаёт целевую величину смещения юнитов
	 * вдоль оси ординат.
	 * <p>
	 * Является обязательным параметром для действия {@code ActionType.ROTATE} и задаёт ординату точки, относительно
	 * которой необходимо совершить поворот.
	 * <p>
	 * Корректными значениями для действия {@code ActionType.MOVE} являются вещественные числа от
	 * {@code -game.WorldHeight} до {@code game.WorldHeight} включительно. Корректными значениями для действия
	 * {@code ActionType.ROTATE} являются вещественные числа от {@code -game.WorldHeight} до
	 * {@code 2.0 * game.WorldHeight} включительно.
	 */
	Y float64
	/**
	 * Задаёт угол поворота.
	 * <p>
	 * Является обязательным параметром для действия {@code ActionType.ROTATE} и задаёт угол поворота относительно точки
	 * ({@code X}, {@code Y}). Положительные значения соответствуют повороту по часовой стрелке.
	 * <p>
	 * Корректными значениями являются вещественные числа от {@code -PI} до {@code PI} включительно.
	 */
	Angle float64
	/**
	 * Устанавливает абсолютное ограничение линейной скорости.
	 * <p>
	 * Является опциональным параметром для действий {@code ActionType.MOVE} и {@code ActionType.ROTATE}. Если для
	 * действия {@code ActionType.ROTATE} установлено ограничение скорости поворота, то этот параметр будет
	 * проигнорирован.
	 * <p>
	 * Корректными значениями являются вещественные неотрицательные числа. При этом, {@code 0.0} означает, что
	 * ограничение отсутствует.
	 */
	MaxSpeed float64
	/**
	 * Устанавливает абсолютное ограничение скорости поворота в радианах за тик.
	 * <p>
	 * Является опциональным параметром для действия {@code ActionType.ROTATE}. Если для этого действия установлено
	 * ограничение скорости поворота, то параметр {@code MaxSpeed} будет проигнорирован.
	 * <p>
	 * Корректными значениями являются вещественные числа в интервале от {@code 0.0} до {@code PI} включительно. При
	 * этом, {@code 0.0} означает, что ограничение отсутствует.
	 */
	MaxAngularSpeed float64
	/**
	 * Устанавливает тип техники.
	 * <p>
	 * Является опциональным параметром для действий {@code ActionType.CLEAR_AND_SELECT},
	 * {@code ActionType.ADD_TO_SELECTION} и {@code ActionType.DESELECT}.
	 * Указанные действия будут применены только к технике выбранного типа.
	 * Параметр будет проигнорирован, если установлена группа юнитов.
	 * <p>
	 * Является опциональным параметром для действия {@code ActionType.SETUP_VEHICLE_PRODUCTION}.
	 * Завод будет настроен на производство техники данного типа. При этом, прогресс производства будет обнулён.
	 * Если данный параметр не установлен, то производство техники на заводе будет остановлено.
	 */
	VehicleType VehicleType
	/**
	 * Устанавливает идентификатор сооружения.
	 * <p>
	 * Является обязательным параметром для действия {@code ActionType.SETUP_VEHICLE_PRODUCTION}.
	 * Если сооружение с данным идентификатором отсутствует в игре, не является заводом по производству техники
	 * ({@code FacilityType.VEHICLE_FACTORY}) или принадлежит другому игроку, то действие будет проигнорировано.
	 */
	FacilityId int64
}

var movesPool sync.Pool

func init() {
	movesPool.New = func() interface{} {
		m := new(Move)
		m.VehicleType = Vehicle_None
		m.Action = Action_None
		return m
	}
}

func NewMove() (m *Move) {
	m = movesPool.Get().(*Move)
	m.Reset()
	return
}

func (m *Move) Reset() {
	m.Action = Action_None
	m.Group = 0
	m.Left = 0
	m.Right = 0
	m.Top = 0
	m.Bottom = 0
	m.X = 0
	m.Y = 0
	m.Angle = 0
	m.MaxSpeed = 0
	m.MaxAngularSpeed = 0
	m.VehicleType = Vehicle_None
	m.FacilityId = 0
}

func (m *Move) Release() {
	movesPool.Put(m)
}

func (m *Move) SelectRect(x, y, width, height float64) {
	m.Left = x
	m.Top = y
	m.Right = x + width
	m.Bottom = y + height
}
